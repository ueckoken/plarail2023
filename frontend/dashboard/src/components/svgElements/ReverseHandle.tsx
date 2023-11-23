// 多分使わない
// 使う場合はindex.tsxのexportに追加

import React from "react";

import { Coordinate } from "@/types/Coordinate";

interface Props {
    coordinate: Coordinate;
    r: number;
    isBack: boolean;
    onChange?: (isBack: boolean) => any;
}

export const ReverseHandle: React.FC<Props> = ({
    coordinate,
    r,
    isBack,
    onChange,
}) => {
    return (
        <>
            <rect
                x={coordinate.x - r * 1.75}
                y={coordinate.y - r * 4}
                width={r * 5}
                height={r * 8}
                fill="black"
            />
            <line
                x1={coordinate.x}
                y1={coordinate.y - r * 1.75}
                x2={coordinate.x}
                y2={coordinate.y + r * 1.75}
                stroke="#555"
                strokeWidth={r * 0.75}
                strokeLinecap="round"
            />
            <circle
                cx={coordinate.x}
                cy={coordinate.y - (isBack ? -1 : 1) * r}
                r={r}
                fill="black"
                stroke="rgba(255, 255, 255, 0.8)"
                strokeWidth={r / 10}
                // onChnageを入れる必要あり？
                onClick={() => {}}
                style={{
                    cursor: "pointer",
                }}
            />
            <g fontSize={r} dominantBaseline="middle" fill="white">
                <text x={coordinate.x + r * 1.5} y={coordinate.y - r * 1.5}>
                    前
                </text>
                <text x={coordinate.x + r * 1.5} y={coordinate.y + r * 1.5}>
                    後
                </text>
            </g>
        </>
    );
};
