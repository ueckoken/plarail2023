import {useQuery} from "@tanstack/react-query";
import {getTrains} from "@/proto/state/v1/state-StateManagerService_connectquery"
import {FC} from "react";

import 'bootstrap/dist/css/bootstrap.min.css'
import {Container, Table} from "react-bootstrap";

type Props = {};

export const TrainList: FC<Props> = ({}) => {
    const {isLoading, isError, error, data} = useQuery({
        ...getTrains.useQuery({}),
        refetchInterval: 200,
    });
    const trains = data?.trains;

    return (
        <Container>
            <h1>Train List</h1>
            {isLoading && <p>読み込み中...</p>}
            {isError && <p role="alert">{error?.message}</p>}
            {!isLoading && !isError && data != null && (
                <div>
                    <Table striped bordered hover>
                        <thead>
                        <tr>
                            <th>Train ID</th>
                            <th>Position ID</th>
                            <th>Priority</th>
                            <th>UUID</th>
                            <th>Destination</th>
                        </tr>
                        </thead>
                        <tbody>
                        {trains?.map((train) => (
                            <tr key={train.trainId}>
                                <td>{train.trainId}</td>
                                <td>{train.positionId}</td>
                                <td>{train.priority}</td>
                                <td>{train.uuid}</td>
                                <td>{train.destination}</td>
                            </tr>
                        ))}
                        </tbody>
                    </Table>
                </div>
            )}
        </Container>
    );
};
