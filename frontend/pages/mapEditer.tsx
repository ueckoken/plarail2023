import React, { FC, useState } from "react";
import Platform, { PlatformProps } from "../components/svgParts/Platform";
import PlatFormComponentForm from "../components/forms/PlatformComponentForm";
import Rail, { RailProps } from "../components/svgParts/Rail";
import StopPoint, { StopPointProps } from "../components/svgParts/StopPoint";
import SwitchPoint, {
    SwitchPointProps,
} from "../components/svgParts/SwitchPoint";

type SVGComponentProps =
    | PlatformProps
    | RailProps
    | StopPointProps
    | SwitchPointProps;
type ComponentType = "platform" | "rail" | "stopPoint" | "switchPoint";

const MapEditerPage: FC<{}> = () => {
    const [mapWidth, setMapWidth] = useState<number>(640);
    const [mapHeight, setMapHeight] = useState<number>(480);
    const [currentComponentType, setCurrentComponentType] =
        useState<ComponentType>("platform");
    const [currentSVGComponents, setCurrentSVGComponents] =
        useState<SVGComponentProps>();
    const [SVGComponents, setSVGComponents] = useState<SVGComponentProps[]>([]);

    return (
        <>
            <fieldset>
                <legend>コンポーネントの追加</legend>
                <div>
                    <label>
                        種類:{" "}
                        <select disabled>
                            <option value="platform">駅</option>
                        </select>
                    </label>
                </div>
                <PlatFormComponentForm
                    onChange={(p: PlatformProps) => {
                        setCurrentSVGComponents(p);
                    }}
                />
                <div>
                    <button
                        onClick={() =>
                            setSVGComponents([
                                ...SVGComponents,
                                currentSVGComponents!,
                            ])
                        }
                    >
                        追加
                    </button>
                </div>
            </fieldset>
            <svg
                width={mapWidth}
                height={mapHeight}
                viewBox={`0 0 ${mapWidth} ${mapHeight}`}
            >
                <rect
                    x={0}
                    y={0}
                    width={mapWidth}
                    height={mapHeight}
                    fill="lightgray"
                />
                {SVGComponents.map((e) => {
                    if ("name" in e)
                        return <Platform name={e.name} position={e.position} />;
                })}
            </svg>
        </>
    );
};

export default MapEditerPage;
