import React from "react";
import clsx from "clsx";

import { Coordinate } from "@/types/Coordinate";
import { degToRad } from "@/utils/convertAngle";

import styles from "./SwitchPoint.module.scss";

interface Props {
    position: Coordinate;
    fromAngle: number;
    leftOutAngle: number;
    rightOutAngle: number;
    isLeft: boolean;
    onClick?: () => void;
}

export const SwitchPoint: React.FC<Props> = ({
    position,
    fromAngle,
    leftOutAngle,
    rightOutAngle,
    isLeft,
    onClick,
}) => {
    const fromAnglePointX = position.x + Math.cos(degToRad(fromAngle)) * 10;
    const fromAnglePointY = position.y + Math.sin(degToRad(fromAngle)) * 10;

    const outAngle = isLeft ? leftOutAngle : rightOutAngle;

    return (
        <g
            onClick={onClick}
            style={{
                cursor: onClick ? "pointer" : "initail",
            }}
        >
            <circle
                cx={position.x}
                cy={position.y}
                r={10}
                className={clsx(
                    styles["switchPointBackground"],
                    styles["switchPointBorder"],
                )}
            />
            <line
                x1={position.x}
                y1={position.y}
                x2={fromAnglePointX}
                y2={fromAnglePointY}
                stroke="black"
                className={clsx(styles["switchPointBorder"])}
            />
            <line
                x1={position.x}
                y1={position.y}
                x2={position.x + 10}
                y2={position.y}
                className={clsx(styles["switchPointBorder"])}
                style={{
                    transform: `rotate(${outAngle}deg)`,
                    transformOrigin: `${position.x}px ${position.y}px`,
                    transition: "transform 0.3s ease-in-out",
                }}
            />
        </g>
    );
};
