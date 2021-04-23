import React from 'react'
import Semaphore, { SelectableStatus } from '../semaphore'
import Row from '../row'

function Table() {
    return (
        <div className="Table">
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th scope="col">Name</th>
                        <th scope="col">Address</th>
                        <th scope="col">Bastion</th>
                        <th scope="col">Local Port</th>
                        <th scope="col">State</th>
                        <th scope="col">Action</th>
                    </tr>
                </thead>
                <tbody>
                    <Row />
                    <tr>
                        <th scope="row">Second</th>
                        <td>192.127.0.1</td>
                        <td>remote.com</td>
                        <td>8080</td>
                        <td><Semaphore status={SelectableStatus.Warning}/></td>
                        <td>-</td>
                    </tr>
                    <tr>
                        <th scope="row">Third</th>
                        <td>10.10.10.101</td>
                        <td>domain.it</td>
                        <td>8090</td>
                        <td><Semaphore status={SelectableStatus.Off}/></td>
                        <td>-</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}

export default Table