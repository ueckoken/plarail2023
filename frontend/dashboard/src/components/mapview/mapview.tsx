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
    addTrain,
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
import {Train as TrainPB} from "@/proto/state/v1/train_pb";

const STOP_RAILS_FROM_ID: Map<string, {x: number, y: number}> = new Map([
  ["shinjyuku_up_s1", { x: 320, y: 400 }],
  ["shinjyuku_down_s1", { x: 200, y: 400 }],
  ["sakurajosui_up_s1", { x: 320, y: 220 }],
  ["sakurajosui_up_s2", { x: 360, y: 220 }],
  ["sakurajosui_down_s1", { x: 200, y: 220 }],
  ["sakurajosui_down_s2", { x: 160, y: 220 }],
  ["chofu_up_s1", { x: 840, y: 220 }],
  ["chofu_up_s2", { x: 800, y: 220 }],
  ["chofu_down_s1", { x: 1000, y: 220 }],
  ["chofu_down_s2", { x: 960, y: 220 }],
  ["hashimoto_up_s1", { x: 700, y: 330 }],
  ["hashimoto_down_s1", { x: 700, y: 410 }],
  ["hachioji_up_s1", { x: 700, y: 480 }],
  ["hachioji_down_s1", { x: 700, y: 560 }],
]);

const POINT_RAILS_FROM_ID: Map<string, {x: number, y: number, fixed: number, straight: number, changed: number}> = new Map([
  ["sakurajosui_up_p1", { x: 320, y: 140, fixed: -90, straight: 90, changed: 45 }],
  ["sakurajosui_down_p1", { x: 200, y: 300, fixed: 90, straight: 270, changed: 225 }], // Use 270 instead of -90, for rotate to clockwise
  ["chofu_down_p1", { x: 1000, y: 140, fixed: -90, straight: 90, changed: 135 }], // Question: これdownじゃない？
]);

const BLOCKS_FROM_ID: Map<string, {x: number, y: number}> = new Map([
  ["shinjyuku_down_b1", { x: 320, y: 400 }],
  ["sakurajosui_down_b1", {x: 320, y: 500}],
  ["chofu_down_b2", {x: 200, y: 300}],
  ["hashimoto_down_b1", {x: 200, y: 140}],
  ["hashimoto_up_b1", {x: 200, y: 220}],
  ["chofu_up_b1", {x: 840, y: 220}],
  ["chofu_down_b1", {x: 1000, y: 220}],
  ["hachioji_down_b1", {x: 700, y: 330}],
  ["hachioji_up_b1", {x: 700, y: 410}],
  ["sakurajosui_up_b1", {x: 700, y: 560}],
  ["shinjyuku_up_b1", {x: 700, y: 560}],
])

const LOCATIONS_FOR_TRAIN_FROM_ID: Map<string, {x: number, y: number}> = new Map([...STOP_RAILS_FROM_ID, ...BLOCKS_FROM_ID]);

function MapView() {
    const { data: stops, refetch: refetchStops } = useSuspenseQuery(getStopStates.useQuery({}));
    const { mutate: mutateStops } = useMutation(updateStopState.useMutation({}));
    const { data: points, refetch: refetchPoints } = useSuspenseQuery(getPointStates.useQuery({}));
    const { mutate: mutatePoints } = useMutation(updatePointState.useMutation({}));
    const { data: trains } = useSuspenseQuery(getTrains.useQuery({}));
    const { mutate: mutateTrain } = useMutation(addTrain.useMutation({}));
    const { data: blocks } = useSuspenseQuery(getBlockStates.useQuery({}));
    const {mutate: mutateBlock} = useMutation(updateBlockState.useMutation({}))

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
    const nextLocation = (train: TrainPB): TrainPB | undefined => {
      const current = train.positionId;
      LOCATIONS_FOR_TRAIN_FROM_ID.forEach((_, key) => {
        if (key === current) {
          const next = LOCATIONS_FOR_TRAIN_FROM_ID.keys().next().value;
          return { ...train, positionId: next };
        }
      });
      return undefined
    }
    return (
        <>
            <button onClick={() => mutateTrain({ train: { trainId: "test", positionId: "shinjyuku_up_s1" }})}>Add Train</button>
            <button onClick={() => mutateTrain({ train: { trainId: "test", positionId: nextLocation(trains.trains[0])?.positionId} })}>Move</button>
            <svg width="100%" height="620px">
                <Background />
                <Rail
                    points={[
                      {x: 320, y: 400},
                      {x: 320, y: 500},
                      {x: 320, y: 500},
                      {x: 200, y: 500},
                      {x: 200, y: 400},
                    ]}
                    isClosed={isClosed("shinjyuku_up_b1")}
                    onClick={() => mutateBlock({state: {blockId: "shinjyuku_up_b1", state: isClosed("shinjyuku_up_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Rail
                    points={[
                      {x: 200, y: 300},
                      {x: 200, y: 400},
                    ]}
                    isClosed={isClosed("shinjyuku_down_b1")}
                    onClick={() => mutateBlock({state: {blockId: "shinjyuku_down_b1", state: isClosed("shinjyuku_down_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Rail
                    points={[
                      {x: 200, y: 300},
                      {x: 200, y: 140},
                    ]}
                />
                <Rail
                    points={[
                      {x: 200, y: 140},
                      {x: 200, y: 30},
                      {x: 200, y: 30},
                      {x: 1000, y: 30},
                      {x: 1000, y: 140},
                    ]}
                    isClosed={isClosed("sakurajosui_down_b1")}
                    onClick={() => mutateBlock({state: {blockId: "sakurajosui_down_b1", state: isClosed("sakurajosui_down_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Rail
                    points={[
                      {x: 1000, y: 140},
                      {x: 1000, y: 220},
                    ]}
                />
                <Rail
                    points={[
                      {x: 1000, y: 220},
                      {x: 1000, y: 560},
                      {x: 700, y: 560},
                    ]}
                    isClosed={isClosed("chofu_down_b1")}
                    onClick={() => mutateBlock({state: {blockId: "chofu_down_b1", state: isClosed("chofu_down_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Rail
                    points={[
                      {x: 700, y: 560},
                      {x: 600, y: 560},
                      {x: 600, y: 480},
                      {x: 700, y: 480},
                    ]}
                    isClosed={isClosed("hachioji_down_b1")}
                    onClick={() => mutateBlock({state: {blockId: "hachioji_down_b1", state: isClosed("hachioji_down_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Rail
                    points={[
                      {x: 700, y: 480},
                      {x: 840, y: 480},
                      {x: 840, y: 220},
                    ]}
                    isClosed={isClosed("hachioji_up_b1")}
                    onClick={() => mutateBlock({state: {blockId: "hachioji_up_b1", state: isClosed("hachioji_up_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Rail
                    points={[
                      {x: 840, y: 250},
                      {x: 840, y: 80},
                      {x: 320, y: 80},
                      {x: 320, y: 140},
                    ]}
                    isClosed={isClosed("chofu_up_b1")}
                    onClick={() => mutateBlock({state: {blockId: "chofu_up_b1", state: isClosed("chofu_up_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Platform isHorizontal={false} position={{ x: 260, y: 220 }} name="桜上水" />
                <Rail
                    points={[
                      {x: 320, y: 140},
                      {x: 320, y: 300},
                    ]}
                />
                <Rail
                    points={[
                      {x: 320, y: 140},
                      {x: 360, y: 170},
                      {x: 360, y: 220},
                      {x: 360, y: 270},
                      {x: 320, y: 300},
                    ]}
                />
                <Rail
                    points={[
                      {x: 200, y: 140},
                      {x: 160, y: 170},
                      {x: 160, y: 220},
                      {x: 160, y: 270},
                      {x: 200, y: 300},
                      ]}
                />
                <Rail
                    points={[
                      {x: 320, y: 300},
                      {x: 320, y: 400},
                    ]}
                    isClosed={isClosed("sakurajosui_up_b1")}
                    onClick={() => mutateBlock({state: {blockId: "sakurajosui_up_b1", state: isClosed("sakurajosui_up_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />
                <Platform isHorizontal={false} position={{ x: 900, y: 220 }} name="調布" />
                <Rail points={[
                    {x: 840, y: 140},
                    {x: 800, y: 170},
                    {x: 800, y: 220},
                ]}></Rail>

                <Rail points={[
                    {x: 800, y: 220},
                    {x: 800, y: 270},
                    {x: 800, y: 330},
                    {x: 700, y: 330},
                ]}
                      isClosed={isClosed("hashimoto_up_b1")}
                      onClick={() => mutateBlock({state: {blockId: "hashimoto_up_b1", state: isClosed("hashimoto_up_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                ></Rail>
                <Rail points={[
                    {x: 700, y: 330},
                    {x: 600, y: 330},
                    {x: 600, y: 410},
                    {x: 700, y: 410},
                ]}
                      isClosed={isClosed("hashimoto_down_b1")}
                        onClick={() => mutateBlock({state: {blockId: "hashimoto_down_b1", state: isClosed("hashimoto_down_b1") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                ></Rail>

                <Rail
                    points={[
                        {x: 960, y: 220},
                        {x: 960, y: 170},
                        {x: 1000, y: 140},
                    ]}
                />

                <Rail
                    points={[
                        {x: 700, y: 410},
                        {x: 800, y: 410},
                        {x: 960, y: 380},
                        {x: 960, y: 220},
                    ]}
                    isClosed={isClosed("chofu_down_b2")}
                    onClick={() => mutateBlock({state: {blockId: "chofu_down_b2", state: isClosed("chofu_down_b2") ? BlockStateEnum.BLOCK_STATE_OPEN : BlockStateEnum.BLOCK_STATE_CLOSE}})}
                />


                <Platform isHorizontal={false} position={{ x: 260, y: 400 }} name="新宿" />
                <Platform position={{ x: 700, y: 370 }} name="橋本" />
                <Platform position={{ x: 700, y: 520 }} name="京王八王子" />
                {
                  stops?.states.map((state) => {
                    const s = STOP_RAILS_FROM_ID.get(state.id)
                    if (s == undefined){
                      throw new Error(`unknown stop rail id: ${state.id}`)
                    }
                    return (
                      <TrafficLight name={state.id} position={s} isStop={state.state == StopStateEnum.STOP_STATE_STOP} key={state.id} onClick={() => mutateStops({state: {id: state.id, state: state.state == StopStateEnum.STOP_STATE_GO ? StopStateEnum.STOP_STATE_STOP : StopStateEnum.STOP_STATE_GO }}, { onSettled: () => refetchStops() })} />
                    )
                  })
                }
                {
                  points?.states.map((state) => {
                    const s = POINT_RAILS_FROM_ID.get(state.id)
                    if (s == undefined){
                      throw new Error(`unknown point rail id: ${state.id}`)
                    }
                    return ( <SwitchPoint position={s} fromAngle={s.fixed} leftOutAngle={s.straight} rightOutAngle={s.changed} isLeft={state.state == PointStateEnum.POINT_STATE_NORMAL} key={state.id} onClick={() => mutatePoints({state: {id: state.id, state: state.state == PointStateEnum.POINT_STATE_NORMAL ? PointStateEnum.POINT_STATE_REVERSE : PointStateEnum.POINT_STATE_NORMAL }}, { onSettled: () => refetchPoints() })} />
                    )
                  })
                }
                {
                  trains?.trains.map((train) => {
                    let s = LOCATIONS_FOR_TRAIN_FROM_ID.get(train.positionId)
                    if (s == undefined){
                      throw new Error(`unknown train location: ${train.positionId}`)
                    }
                    s = {x: s.x, y: s.y+30}
                    return ( <Train position={s} key={train.trainId} trainData={{id: train.trainId, type: train.priority}}/>
                    )
                  })

                }

            </svg>
        </>
    );
}

export default MapView
