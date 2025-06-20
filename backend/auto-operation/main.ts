import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";
import { StateManagerService } from "./proto/state/v1/state_connect.js";
import { BlockStateEnum } from "./proto/state/v1/block_pb.js";
import { StopStateEnum } from "./proto/state/v1/stop_pb.js";
import { PointStateEnum } from "./proto/state/v1/point_pb.js";
import { Priority } from "./proto/state/v1/train_pb.js";
// @ts-ignore
import config from './map/mfk-2024.json' assert {type: 'json'};

type MapConfig = {
  stopBlocks: { [key: string]: string },
  stations: {
    [key: string]: {
      capacity: number,
    }
  }
}
let mapConfig: MapConfig;

const SERVER_ADDR = process.env['SERVER_ADDR'] ?? 'http://localhost:8080'
const transport = createConnectTransport(
  {
    httpVersion: "2",
    baseUrl: SERVER_ADDR,
  });
const client = createPromiseClient(StateManagerService, transport);

async function loadConfig() {
  // @ts-ignore
  mapConfig = config;
}

loadConfig();

async function addTest() {
  // テスト用の列車を追加
  await client.addTrain({
    train: {
      trainId: "test",
      positionId: "shinjuku_b1",
      priority: Priority.HIGH,
      uuid: "test",
      destination: "hashimoto_up_s1"
    }
  })
  await client.addTrain({
    train: {
      trainId: "test2",
      positionId: "shinjuku_s1",
      priority: Priority.LOW,
      uuid: "test",
      destination: "hachioji_up_s1"
    }
  })
}

async function main() {

  // 閉塞確認
  const blocks = (await client.getBlockStates({})).states;
  const stops = (await client.getStopStates({})).states;
  const points = (await client.getPointStates({})).states;
  const trains = (await client.getTrains({})).trains;
  for (const stop of stops) {
    const collespondBlock = mapConfig.stopBlocks[stop.id];
    // 閉塞の確認する
    const block = blocks.find(b => b.blockId === collespondBlock);
    if (block && block.state === BlockStateEnum.BLOCK_STATE_CLOSE) {
      if (stop.state !== StopStateEnum.STOP_STATE_STOP) {
        // 閉塞が閉じていたら列車を止めておく
        await client.updateStopState({
          "state": {
            "id": stop.id,
            "state": StopStateEnum.STOP_STATE_STOP
          }
        })
      }
    } else {
      if (stop.state !== StopStateEnum.STOP_STATE_GO) {
        // 閉塞が開いていたら列車を開けておく
        await client.updateStopState({
          "state": {
            "id": stop.id,
            "state": StopStateEnum.STOP_STATE_GO
          }
        })
      }
    }
  }

  // 列車が発車して良いかの確認
  for (const train of trains) {
    // 列車が駅に停車しているかどうか
    if ((new RegExp(".*_.*_s[0-9]")).test(train.positionId)) {
      // 駅に停車している場合は進行して良いかの確認をする
      const collespondBlock = mapConfig.stopBlocks[train.positionId];
      // 対応する閉塞が空いているかを確認する
      const block = blocks.find(b => b.blockId === collespondBlock);
      // 閉塞が閉まっていたら処理を打ち切る
      if (block && block.state !== BlockStateEnum.BLOCK_STATE_OPEN) continue;
      // 列車のPriorityによって閉塞を開けるかどうかを決める
      /*
        通過待ちができる駅で、PriorityがLOWの列車はPriorityがHIGHの列車が停車するまで待つ
      */
      // 通過待ちができるかどうか
      const stationName = train.positionId.split("_")[0] + "_" + train.positionId.split("_")[1];
      const capacity = mapConfig.stations[stationName];
      if (capacity.capacity > 1 && train.priority === Priority.LOW) {
        // 通過待ちが可能な駅で、PriorityがLOWの列車はPriorityがHIGHの列車が停車するまで待つ
        const highPriorityTrains = trains.filter(t => (t.positionId.includes(stationName)) && t.priority === Priority.HIGH);
        if (highPriorityTrains.length > 0) continue;
      }
      const stop = stops.find(s => s.id === train.positionId);
      // 問題ないならGOにする
      if (stop && stop.state !== StopStateEnum.STOP_STATE_GO) {
        await client.updateStopState({
          "state": {
            "id": stop.id,
            "state": StopStateEnum.STOP_STATE_GO
          }
        })
      }
    }
  }


  // for (const point of points) {
  //   if (point.id == "ninini_p1") {
  //     // ninini_b1のblockが閉じたらninini_p1を逆にする
  //     const ninini_b1 = blocks.find(b => b.blockId === "ninini_b1");
  //     if (ninini_b1 && ninini_b1.state === BlockStateEnum.BLOCK_STATE_CLOSE) {
  //       if (point.state !== PointStateEnum.POINT_STATE_REVERSE) {
  //         await client.updatePointState({
  //           "state": {
  //             "id": point.id,
  //             "state": PointStateEnum.POINT_STATE_REVERSE
  //           }
  //         })
  //       }
  //     } else {
  //       if (point.state !== PointStateEnum.POINT_STATE_NORMAL) {
  //         await client.updatePointState({
  //           "state": {
  //             "id": point.id,
  //             "state": PointStateEnum.POINT_STATE_NORMAL
  //           }
  //         })
  //       }
  //     }
  //   }
  // }
  // for (const point of points) {
  //   console.log(point);
  // }

  // ポイント確認
  // for (const point of points) {
  //   if (point.id === "sakurajosui_up_p1") {
  //     // 桜上水上りポイント
  //     // デフォルトではSTRAIGHTにして、sakurajosui_up_s1がONならREVERSEにする
  //     const sakurajosui_up_s1 = stops.find(s => s.id === "sakurajosui_up_s1");
  //     if (sakurajosui_up_s1 && sakurajosui_up_s1.state === StopStateEnum.STOP_STATE_STOP) {
  //       if (point.state !== PointStateEnum.POINT_STATE_REVERSE) {
  //         await client.updatePointState({
  //           "state": {
  //             "id": point.id,
  //             "state": PointStateEnum.POINT_STATE_REVERSE
  //           }
  //         })
  //       }
  //     } else {
  //       if (point.state !== PointStateEnum.POINT_STATE_NORMAL) {
  //         await client.updatePointState({
  //           "state": {
  //             "id": point.id,
  //             "state": PointStateEnum.POINT_STATE_NORMAL
  //           }
  //         })
  //       }
  //     }
  //   }
  //   if (point.id === "sakurajosui_down_p1") {
  //     // const sakurajosui_down_s1 = stops.find(s => s.id === "sakurajosui_down_s1");
  //     const sakurajosui_down_b1 = blocks.find(b => b.blockId === "sakurajosui_down_b1");
  //     if (sakurajosui_down_b1 && sakurajosui_down_b1.state === BlockStateEnum.BLOCK_STATE_CLOSE) {
  //       if (point.state !== PointStateEnum.POINT_STATE_REVERSE) {
  //         await client.updatePointState({
  //           "state": {
  //             "id": point.id,
  //             "state": PointStateEnum.POINT_STATE_REVERSE
  //           }
  //         })
  //       }
  //     } else {
  //       if (point.state !== PointStateEnum.POINT_STATE_NORMAL) {
  //         await client.updatePointState({
  //           "state": {
  //             "id": point.id,
  //             "state": PointStateEnum.POINT_STATE_NORMAL
  //           }
  //         })
  //       }
  //     }
  //   }
  //   if (point.id === "chofu_up_p1") {
  //     // 閉塞が空いてる方に切り替える b1 or b2
  //     const chofu_up_b1 = blocks.find(b => b.blockId === "chofu_up_b1");
  //     const chofu_up_b2 = blocks.find(b => b.blockId === "chofu_up_b2");
  //     if (chofu_up_b1 && chofu_up_b2) {
  //       if (chofu_up_b1.state === BlockStateEnum.BLOCK_STATE_OPEN) {
  //         if (point.state !== PointStateEnum.POINT_STATE_NORMAL) {
  //           await client.updatePointState({
  //             "state": {
  //               "id": point.id,
  //               "state": PointStateEnum.POINT_STATE_NORMAL
  //             }
  //           })
  //         }
  //       } else if (chofu_up_b2.state === BlockStateEnum.BLOCK_STATE_OPEN) {
  //         if (point.state !== PointStateEnum.POINT_STATE_REVERSE) {
  //           await client.updatePointState({
  //             "state": {
  //               "id": point.id,
  //               "state": PointStateEnum.POINT_STATE_REVERSE
  //             }
  //           })
  //         }
  //       }
  //     }
  //   }

  //   //   // 調布ポイント
  //   //   // 閉塞 chofu_up_b1 にいる列車の目的地に応じてポイントを切り替える
  //   //   const train = trains.find(t => t.positionId === "chofu_up_b1");
  //   //   if (train) {
  //   //     const destination = train.destination;
  //   //     if (destination.includes("hashimoto")) {
  //   //       if (point.state !== PointStateEnum.POINT_STATE_REVERSE) {
  //   //         await client.updatePointState({
  //   //           "state": {
  //   //             "id": point.id,
  //   //             "state": PointStateEnum.POINT_STATE_REVERSE
  //   //           }
  //   //         })
  //   //       }
  //   //     } else {
  //   //       if (point.state !== PointStateEnum.POINT_STATE_NORMAL) {
  //   //         await client.updatePointState({
  //   //           "state": {
  //   //             "id": point.id,
  //   //             "state": PointStateEnum.POINT_STATE_NORMAL
  //   //           }
  //   //         })
  //   //       }

  //   //     }
  //   //   }
  //   // }
  // }
}

// addTest();

(async () => {
  while (true) {
    main();
    await new Promise(resolve => setTimeout(resolve, 200));
  }
})()
