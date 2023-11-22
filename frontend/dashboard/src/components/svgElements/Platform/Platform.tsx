import React from "react";
import clsx from "clsx";

import { Coordinate } from "@/types/Coordinate";

import styles from "./Platform.module.scss";

interface Props {
    name: string;
    position: Coordinate;
    isHorizontal?: boolean;
    size?: number;
}

export const Platform: React.FC<Props> = ({
    name,
    position,
    isHorizontal = true,
    size,
}) => {
    const stationSize = size ? size : 20 * name.length + 10;

    const width = isHorizontal ? stationSize : 30;
    const height = isHorizontal ? 30 : stationSize;

    return (
        <>
            <g>
                <rect
                    x={position.x - width / 2}
                    y={position.y - height / 2}
                    width={width}
                    height={height}
                    className={clsx(styles["platform"])}
                />
                <text
                    x={position.x}
                    y={position.y}
                    width={width}
                    height={height}
                    fontFamily="monospace"
                    fontSize={20}
                    writingMode={isHorizontal ? "horizontal-tb" : "vertical-lr"}
                    textAnchor="middle"
                    dominantBaseline="central"
                >
                    {name}
                </text>
            </g>
        </>
    );
};
