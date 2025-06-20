import React, { FC, useState } from "react";
import { Rect, Stage, Layer } from "react-konva";
import { Platform, PlatformProps } from "../components/MapParts";
import PlatFormComponentForm from "../components/forms/PlatformComponentForm";

type SVGComponentProps = PlatformProps;
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
<div>
                <h3>駅リスト</h3>
                {SVGComponents.map((e, idx) => (
                    <div key={idx} style={{ display: "flex", alignItems: "center", marginBottom: 4 }}>
                        <span style={{ marginRight: 8 }}>{e.name}</span>
                        <button onClick={() => setCurrentSVGComponents(e)}>編集</button>
                    </div>
                ))}
            </div>
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
            <Stage width={mapWidth} height={mapHeight}>
                <Layer>
                    <Rect
                        x={0}
                        y={0}
                        width={mapWidth}
                        height={mapHeight}
                        fill="lightgray"
                    />
                    {SVGComponents.map((e) => {
                        if ("name" in e)
                            return (
                                <Platform name={e.name} position={e.position} />
                            );
                    })}
                </Layer>
            </Stage>
        </>
    );
};

export default MapEditerPage;
