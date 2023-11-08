import { FC } from "react";
import { Stage, Group, Rect, Text } from "react-konva";
import { Point } from "../types/svgPartsTypes";

export interface PlatformProps {
    name: string;
    position: Point;
    isHorizontal?: boolean;
    size?: number;
}
const Platform: FC<PlatformProps> = ({
    name,
    position,
    isHorizontal = true,
    size = 40,
}) => {
    const width = isHorizontal ? size : 20;
    const height = isHorizontal ? 20 : size;
    return (
        <Group draggable>
            <Rect
                x={position.x - width / 2}
                y={position.y - height / 2}
                width={width}
                height={height}
                fill="white"
                stroke="black"
            />
            <Text
                text={name}
                x={position.x}
                y={position.y}
                align="center"
                verticalAlign="middle"
                fontSize={20}
            />
        </Group>
    );
};

export { Platform };
