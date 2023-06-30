import { FC, useEffect, useState } from "react";
import { PlatformProps } from "../svgParts/Platform";

interface Props {
    onChange: (p: PlatformProps) => void;
}

const PlatFormComponentForm: FC<Props> = ({ onChange }) => {
    const [name, setName] = useState<string>("");
    const [x, setX] = useState<number>(0);
    const [y, setY] = useState<number>(0);

    useEffect(() => {
        onChange({
            name,
            position: { x, y },
        });
    }, [name, x, y]);

    return (
        <>
            <div>
                <label>
                    駅名:{" "}
                    <input
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                </label>
            </div>
            <div>
                <label>
                    x:{" "}
                    <input
                        type="number"
                        value={x}
                        onChange={(e) => setX(+e.target.value)}
                    />
                </label>
            </div>
            <div>
                <label>
                    y:{" "}
                    <input
                        type="number"
                        value={y}
                        onChange={(e) => setY(+e.target.value)}
                    />
                </label>
            </div>
        </>
    );
};

export default PlatFormComponentForm;
