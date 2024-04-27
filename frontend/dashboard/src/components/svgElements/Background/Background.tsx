import React from "react";
import clsx from "clsx";

import styles from "./Background.module.scss";

export const Background: React.FC = ({ }) => {
    return (
        <>
            <g>
                <rect
                    x={0}
                    y={0}
                    width={700}
                    height={400}
                    className={clsx(styles["background"])}
                />
            </g>
        </>
    );
};
