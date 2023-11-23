import React from "react";
import clsx from "clsx";

import { Coordinate } from "@/types/Coordinate";

import styles from "./Rail.module.scss";

interface Props {
    points: [Coordinate, Coordinate, ...Coordinate[]];
    isClosed?: boolean;
}

export const Rail: React.FC<Props> = ({ points, isClosed }) => {
    const pointsText = points
        .map((point: Coordinate) => `${point.x}, ${point.y}`)
        .join("\n");

    return (
        <>
            <g>
                <polyline
                    points={pointsText}
                    fill="none"
                    className={clsx(
                        styles[isClosed ? "closedRail" : "openRail"],
                    )}
                />
            </g>
        </>
    );
};
