import React from "react";
import clsx from "clsx";

import { Coordinate } from "@/types/Coordinate";

import styles from "./TrafficLight.module.scss";

interface Props {
    position: Coordinate;
    name?: string;
    isStop: boolean;
    onClick?: () => void;
}

export const TrafficLight: React.FC<Props> = ({ position, name, isStop, onClick }) => {
    return (
        <g
            onClick={onClick}
            style={{
                cursor: onClick ? "pointer" : "initial",
            }}
        >
            {
              name && (
                <text
                x={position.x}
                y={position.y - 20}
                fontSize={8}
                >{name}</text>
                )
            }
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
