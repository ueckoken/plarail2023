import {useQuery} from "@tanstack/react-query";
import {getTrains} from "@/proto/state/v1/state-StateManagerService_connectquery"
import {FC} from "react";

type Props = {
    userId: string;
};

export const TrainList: FC<Props> = ({userId}) => {
    const {isLoading, isError, error, data} = useQuery(
        getTrains.useQuery({userId})
    );

    return (
        <div>
            {isLoading && <p>読み込み中...</p>}
            {isError && <p role="alert">{error?.message}</p>}
            {!isLoading && !isError && data != null && <h3>{data?.trains.length}</h3>}
        </div>
    );
};
