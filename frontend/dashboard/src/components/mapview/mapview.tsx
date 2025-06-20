import {
  Rail,
  Platform,
  TrafficLight,
  SwitchPoint,
  Background,
  Train,
} from "@/components/svgElements";
import { useMutation, useSuspenseQuery } from '@tanstack/react-query';
import {
  // addTrain,
  getBlockStates,
  getPointStates,
  getStopStates,
  getTrains,
  updateBlockState,
  updatePointState,
  updateStopState
} from '@/proto/state/v1/state-StateManagerService_connectquery'
import { StopStateEnum } from "@/proto/state/v1/stop_pb";
import { PointStateEnum } from "@/proto/state/v1/point_pb";
import { BlockState, BlockStateEnum } from "@/proto/state/v1/block_pb";
// import { Train as TrainPB } from "@/proto/state/v1/train_pb";

const STOP_RAILS_FROM_ID: Map<string, { x: number, y: number }> = new Map([
  ["ninini_s1", { x: 350, y: 100 }],
  ["ninini_s2", { x: 350, y: 150 }],
  ["nanana_s1", { x: 350, y: 300 }],
]);

const POINT_RAILS_FROM_ID: Map<string, { x: number, y: number, fixed: number, straight: number, changed: number }> = new Map([
  ["ninini_p1", { x: 500, y: 100, fixed: 0, straight: 180, changed: 140 }],
]);

const BLOCKS_FROM_ID: Map<string, { x: number, y: number }> = new Map([
  ["ninini_b1", { x: 320, y: 400 }],
  ["nanana_b1", { x: 320, y: 500 }],
])

const LOCATIONS_FOR_TRAIN_FROM_ID: Map<string, { x: number, y: number }> = new Map([...STOP_RAILS_FROM_ID, ...BLOCKS_FROM_ID]);

function MapView() {
  const { data: stops, refetch: refetchStops } = useSuspenseQuery(getStopStates.useQuery({}));
  const { mutate: mutateStops } = useMutation(updateStopState.useMutation({}));
  const { data: points, refetch: refetchPoints } = useSuspenseQuery(getPointStates.useQuery({}));
  const { mutate: mutatePoints } = useMutation(updatePointState.useMutation({}));
  const { data: trains } = useSuspenseQuery(getTrains.useQuery({}));
  // const { mutate: mutateTrain } = useMutation(addTrain.useMutation({}));
  const { data: blocks } = useSuspenseQuery(getBlockStates.useQuery({}));
  const { mutate: mutateBlock } = useMutation(updateBlockState.useMutation({}))

  // blocks to Map
  const blocksMap = new Map<string, BlockState>();
  blocks?.states.forEach((block) => {
    blocksMap.set(block.blockId, block);
  });

  function isClosed(blockId: string): boolean {
    const s = blocksMap.get(blockId)?.state;
    console.log(blockId, s);
    if (s == undefined) {
      console.log(`blockId ${blockId} is not found`);
    }
    return s === BlockStateEnum.BLOCK_STATE_CLOSE;
  }
  // const nextLocation = (train: TrainPB): TrainPB | undefined => {
  //   const current = train.positionId;
  //   LOCATIONS_FOR_TRAIN_FROM_ID.forEach((_, key) => {
  //     if (key === current) {
  //       const next = LOCATIONS_FOR_TRAIN_FROM_ID.keys().next().value;
  //       return { ...train, positionId: next };
  //     }
  //   });
  //   return undefined
  // }
  return (
    <div style={{ width: '100%', height: '100%', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      {/* <button onClick={() => mutateTrain({ train: { trainId: "test", positionId: "shinjyuku_up_s1" } })}>Add Train</button>
      <button onClick={() => mutateTrain({ train: { trainId: "test", positionId: nextLocation(trains.trains[0])?.positionId } })}>Move</button> */}
      <svg width="100%" height="100%" viewBox="0 0 700 400" preserveAspectRatio="xMidYMid meet" style={{ display: 'block', maxWidth: '100%', maxHeight: '100%' }}>
        <Background />
        <Rail
          points={[
            { x: 200, y: 100 },
            { x: 100, y: 100 },
            { x: 100, y: 300 },
            { x: 200, y: 300 },
            { x: 350, y: 300 },
          ]}
          isClosed={isClosed("nanana_b1")}
          onClick={() => mutateBlock({ state: { blockId: "nanana_b1", state: isClosed("nanana_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE } })}
        />
        <Rail
          points={[
            { x: 350, y: 300 },
            { x: 600, y: 300 },
            { x: 600, y: 100 },
            { x: 500, y: 100 }
          ]}
          isClosed={isClosed("ninini_b1")}
          onClick={() => mutateBlock({ state: { blockId: "ninini_b1", state: isClosed("ninini_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE } })}
        />

        <Rail
          points={[
            { x: 200, y: 100 },
            { x: 250, y: 150 },
            { x: 250, y: 150 },
            { x: 450, y: 150 },
            { x: 450, y: 150 },
            { x: 500, y: 100 },

            { x: 200, y: 100 },
            { x: 500, y: 100 }
          ]}
        />

        <Platform isHorizontal={true} position={{ x: 350, y: 50 }} name="ニニニ" />
        <Platform isHorizontal={true} position={{ x: 350, y: 350 }} name="ナナナ" />
        {
          stops?.states.map((state) => {
            const s = STOP_RAILS_FROM_ID.get(state.id)
            if (s == undefined) {
              throw new Error(`unknown stop rail id: ${state.id}`)
            }
            return (
              <TrafficLight name={state.id} position={s} isStop={state.state == StopStateEnum.STOP_STATE_STOP} key={state.id} onClick={() => mutateStops({ state: { id: state.id, state: state.state == StopStateEnum.STOP_STATE_GO ? StopStateEnum.STOP_STATE_STOP : StopStateEnum.STOP_STATE_GO } }, { onSettled: () => refetchStops() })} />
            )
          })
        }
        {
          points?.states.map((state) => {
            const s = POINT_RAILS_FROM_ID.get(state.id)
            if (s == undefined) {
              throw new Error(`unknown point rail id: ${state.id}`)
            }
            return (<SwitchPoint position={s} fromAngle={s.fixed} leftOutAngle={s.straight} rightOutAngle={s.changed} isLeft={state.state == PointStateEnum.POINT_STATE_NORMAL} key={state.id} onClick={() => mutatePoints({ state: { id: state.id, state: state.state == PointStateEnum.POINT_STATE_NORMAL ? PointStateEnum.POINT_STATE_REVERSE : PointStateEnum.POINT_STATE_NORMAL } }, { onSettled: () => refetchPoints() })} />
            )
          })
        }
        {
          trains?.trains.map((train) => {
            let s = LOCATIONS_FOR_TRAIN_FROM_ID.get(train.positionId)
            if (s == undefined) {
              throw new Error(`unknown train location: ${train.positionId}`)
            }
            s = { x: s.x, y: s.y + 30 }
            return (<Train position={s} key={train.trainId} trainData={{ id: train.trainId, type: train.priority }} />
            )
          })

        }

      </svg>
    </div>
  );
}

export default MapView
