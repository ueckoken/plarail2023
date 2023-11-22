import React from "react";
import clsx from "clsx";

import { Coordinate } from "@/types/Coordinate";

import styles from "./TraficLight.module.scss";

interface Props {
    position: Coordinate;
    isStop: boolean;
    onClick?: () => void;
}

export const TraficLight: React.FC<Props> = ({ position, isStop, onClick }) => {
    return (
        <g
            onClick={onClick}
            style={{
                cursor: onClick ? "pointer" : "initial",
            }}
        >
            <line
                x1={position.x}
                y1={position.y - 5}
                x2={position.x}
                y2={position.y + 5}
                strokeWidth={20}
                strokeLinecap="round"
                className={clsx(styles["traficLightBackground"])}
            />
            <circle
                cx={position.x}
                cy={position.y - 6}
                r={5}
                className={clsx(
                    styles[isStop ? "traficLightOff" : "traficLightGo"],
                )}
            />
            <circle
                cx={position.x}
                cy={position.y + 6}
                r={5}
                className={clsx(
                    styles[isStop ? "traficLightStop" : "traficLightOff"],
                )}
            />
        </g>
    );
};
