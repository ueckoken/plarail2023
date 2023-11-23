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
import { getStopStates, updateStopState } from '@/proto/state/v1/state-StateManagerService_connectquery'
import { StopStateEnum } from "@/proto/state/v1/stop_pb";

type StopRailPosition = {
  id: string,
  position: {x: number, y: number},
}

const STOP_RAILS_FROM_ID: Map<string, {x: number, y: number}> = new Map([
  ["shinjyuku_up_s1", { x: 700, y: 40 }],
  ["shinjyuku_down_s1", { x: 700, y: 100 }],
  ["sakurajosui_up_s1", { x: 950, y: 250 }],
  ["sakurajosui_up_s2", { x: 1000, y: 400 }],
  ["sakurajosui_down_s1", { x: 950, y: 400 }],
  ["sakurajosui_down_s2", { x: 900, y: 400 }],
  ["chofu_up_s1", { x: 350, y: 550 }],
  ["chofu_up_s2", { x: 350, y: 400 }],
  ["chofu_down_s1", { x: 400, y: 400 }],
  ["chofu_down_s2", { x: 300, y: 400 }],
  ["hashimoto_up_s1", { x: 90, y: 300 }],
  ["hashimoto_down_s1", { x: 90, y: 240 }],
  ["hachioji_up_s1", { x: 90, y: 40 }],
  ["hachioji_down_s1", { x: 90, y: 100 }],
])

export default function Home() {
    const { data } = useQuery(getStopStates.useQuery({}));
    const { mutate } = useMutation(updateStopState.useMutation({}));
    return (
        <>
            <svg width="100%" height="620px">
                <Background />
                <Rail
                    points={[
                        { x: 50, y: 100 },
                        { x: 250, y: 100 },
                        { x: 250, y: 200 },
                        { x: 50, y: 200 },
                        { x: 50, y: 100 },
                    ]}
                />
                <SwitchPoint
                    position={{ x: 250, y: 150 }}
                    fromAngle={0}
                    leftOutAngle={90}
                    rightOutAngle={-90}
                    isLeft={true}
                />
                <Platform position={{ x: 150, y: 150 }} name="新線新宿" />
                {
                  data?.states.map((state) => {
                    const s = STOP_RAILS_FROM_ID.get(state.id)
                    if (s == undefined){
                      throw new Error(`unknown stop rail id: ${state.id}`)
                    }
                    return (
                      <TrafficLight position={s} isStop={state.state == StopStateEnum.STOP_STATE_STOP} key={state.id} onClick={() => mutate({state: {id: state.id, state: state.state == StopStateEnum.STOP_STATE_GO ? StopStateEnum.STOP_STATE_STOP : StopStateEnum.STOP_STATE_GO }})} />
                    )
                  })
                }
                <Train
                    position={{ x: 150, y: 100 }}
                    trainData={{ id: "keio", type: 2 }}
                />
                <Train
                    position={{ x: 150, y: 200 }}
                    trainData={{ id: "shinjyuku", type: 1 }}
                />
            </svg>
        </>
    );
}
