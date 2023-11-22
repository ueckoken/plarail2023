import React from "react";

import clsx from "clsx";

import styles from "./Train.module.scss";

import { Coordinate } from "@/types/Coordinate";
import { TrainData } from "@/types/TrainData";

import TrainIcon from "@/../public/train.svg";

interface Props {
    position: Coordinate;
    trainData: TrainData;
}

export const Train: React.FC<Props> = ({ position, trainData }) => {
    const trainTypeList: [string, string][] = [
        ["各停", "black"],
        ["急行", "limegreen"],
        ["特急", "red"],
    ];
    const trainType: [string, string] =
        trainTypeList[trainData.type ? trainData.type : 0];

    const trainIconWidth: number = 40;
    const trainIconHeight: number = 40;

    const trainTypeTextSize: number = 13;

    return (
        <>
            <g>
                <TrainIcon
                    x={position.x - trainIconWidth / 2}
                    y={position.y - trainIconHeight / 2}
                    width={trainIconHeight}
                    height={trainIconHeight}
                />
                <rect
                    x={position.x - (trainIconWidth - 10) / 2}
                    y={
                        position.y -
                        trainTypeTextSize / 2 -
                        trainIconHeight +
                        15
                    }
                    width={trainIconWidth - 10}
                    height={trainTypeTextSize}
                    fill={trainType[1]}
                />
                <text
                    x={position.x}
                    y={position.y - trainIconHeight + 15}
                    width={trainIconWidth}
                    height={trainTypeTextSize}
                    fontFamily="monospace"
                    fontSize={trainTypeTextSize}
                    textAnchor="middle"
                    dominantBaseline="central"
                    fill="white"
                    fontWeight={"bold"}
                >
                    {trainType[0]}
                </text>
                <text
                    x={position.x}
                    y={position.y + trainIconHeight - 15}
                    width={40}
                    height={16}
                    fontFamily="monospace"
                    fontSize={10}
                    textAnchor="middle"
                    dominantBaseline="central"
                    fill="black"
                    fontWeight={"bold"}
                >
                    {trainData.id}
                </text>
            </g>
        </>
    );
};
