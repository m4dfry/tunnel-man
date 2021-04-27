import React, { FunctionComponent } from 'react'
import Semaphore, { SelectableStatus } from '../semaphore'
import { TunnelState } from '../../store/tunnels'
import Row from '../row'

interface TableProps {
    tunnels: Array<TunnelState>
}

export const Table: FunctionComponent<TableProps> = (
        props: TableProps
    ) => {

    return (


        <div className="Table">
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th scope="col">Name</th>
                        <th scope="col">Bastion</th>
                        <th scope="col">Address</th>
                        <th scope="col">Local Port</th>
                        <th scope="col">State</th>
                        <th scope="col">Action</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        props.tunnels.map((tunnel: TunnelState, i) => {
                            return (
                                <Row tunnel={tunnel} key={i} />        
                            );
                        })
                    }
                </tbody>
            </table>
        </div>
    )
}

export default Table;