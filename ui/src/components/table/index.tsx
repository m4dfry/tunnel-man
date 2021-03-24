import React, { useState } from 'react'
import Semaphore, { SelectableStatus } from '../semaphore'

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
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <th scope="row">First</th>
                        <td>127.0.0.1</td>
                        <td>Local</td>
                        <td>80</td>
                        <td><Semaphore status={SelectableStatus.Running}/></td>
                    </tr>
                    <tr>
                        <th scope="row">Second</th>
                        <td>192.127.0.1</td>
                        <td>remote.com</td>
                        <td>8080</td>
                        <td><Semaphore status={SelectableStatus.Warning}/></td>
                    </tr>
                    <tr>
                        <th scope="row">Third</th>
                        <td>10.10.10.101</td>
                        <td>domain.it</td>
                        <td>8090</td>
                        <td><Semaphore status={SelectableStatus.Off}/></td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}

export default Table