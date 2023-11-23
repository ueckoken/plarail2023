import {
    Rail,
    Platform,
    TraficLight,
    SwitchPoint,
    Background,
    Train,
} from "@/components/svgElements";

export default function Home() {
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
                <TraficLight position={{ x: 150, y: 200 }} isStop={false} />
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
