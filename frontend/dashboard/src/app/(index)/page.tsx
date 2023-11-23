"use client";

import {
    Rail,
    Platform,
    TrafficLight,
    SwitchPoint,
    Background,
    Train,
} from "@/components/svgElements";
import { useMutation, useQuery } from '@tanstack/react-query';
import { getPointStates, getStopStates, getTrains, updatePointState, updateStopState } from '@/proto/state/v1/state-StateManagerService_connectquery'
import { StopStateEnum } from "@/proto/state/v1/stop_pb";
import {PointStateEnum} from "@/proto/state/v1/point_pb";

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
  ["chofu_up_p1", { x: 1000, y: 140, fixed: -90, straight: 90, changed: 135 }], // Question: これdownじゃない？
]);

const LOCATIONS_FOR_TRAIN_FROM_ID: Map<string, {x: number, y: number}> = new Map([...STOP_RAILS_FROM_ID, ...POINT_RAILS_FROM_ID])

export default function Home() {
    const { data: stops, refetch: refetchStops } = useQuery(getStopStates.useQuery({}));
    const { mutate: mutateStops } = useMutation(updateStopState.useMutation({}));
    const { data: points, refetch: refetchPoints } = useQuery(getPointStates.useQuery({}));
    const { mutate: mutatePoints } = useMutation(updatePointState.useMutation({}));
    const { data: trains } = useQuery(getTrains.useQuery({}));
    return (
        <>
            <svg width="100%" height="620px">
                <Background />
                <Rail
                    points={[
                      {x: 320, y: 400},
                      {x: 320, y: 500},
                      {x: 320, y: 500},
                      {x: 200, y: 500},
                      {x: 200, y: 400},
                      {x: 200, y: 30},
                      {x: 1000, y: 30},
                      {x: 1000, y: 180},
                      {x: 1000, y: 560},
                      {x: 700, y: 560},
                      {x: 600, y: 560},
                      {x: 600, y: 480},
                      {x: 700, y: 480},
                      {x: 840, y: 480},
                      {x: 840, y: 180},
                      {x: 840, y: 80},
                      {x: 320, y: 80},
                      {x: 320, y: 180},
                      {x: 320, y: 400},
                    ]}
                />
                <Platform isHorizontal={false} position={{ x: 260, y: 220 }} name="桜上水" />
                <Rail
                    points={[
                      {x: 320, y: 140},
                      {x: 360, y: 170},
                      {x: 360, y: 220},
                      {x: 360, y: 270},
                      {x: 320, y: 300},
                      {x: 320, y: 400},
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
                <Platform isHorizontal={false} position={{ x: 900, y: 220 }} name="調布" />
                <Rail
                    points={[
                      {x: 840, y: 140},
                      {x: 800, y: 170},
                      {x: 800, y: 220},
                      {x: 800, y: 270},
                      {x: 800, y: 330},
                      {x: 700, y: 330},
                      {x: 600, y: 330},
                      {x: 600, y: 410},
                      {x: 700, y: 410},
                      {x: 800, y: 410},
                      {x: 960, y: 380},
                      {x: 960, y: 220},
                      {x: 960, y: 170},
                      {x: 1000, y: 140},
                    ]}
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
                    const s = LOCATIONS_FOR_TRAIN_FROM_ID.get(train.positionId)
                    if (s == undefined){
                      throw new Error(`unknown train location: ${train.positionId}`)
                    }
                    return ( <Train position={s} key={train.trainId} trainData={{id: train.trainId, type: train.priority}}/>
                    )
                  })

                }

            </svg>
        </>
    );
}
