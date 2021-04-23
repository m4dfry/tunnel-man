import React from 'react'
import './index.css'
import Semaphore, { SelectableStatus } from '../semaphore'
import Action from '../action'

function Row() {
    return (
        <tr>
            <th scope="row">First</th>
            <td>127.0.0.1</td>
            <td>Local</td>
            <td>80</td>
            <td><Semaphore status={SelectableStatus.Running}/></td>
            <td><Action /></td>
        </tr> 
    )
}

export default Row