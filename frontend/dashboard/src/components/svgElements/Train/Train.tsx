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
    return (
        <>
            <TrainIcon x={position.x} y={position.y} width={50} height={50} />
        </>
    );
};
