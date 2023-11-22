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
                <Platform position={{ x: 150, y: 50 }} name="駅" />
                <Train position={{ x: 50, y: 50 }} trainData={{ id: "id" }} />
            </svg>
        </>
    );
}
