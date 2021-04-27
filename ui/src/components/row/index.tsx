import React, { FunctionComponent } from 'react'
import './index.css'
import { TunnelState } from '../../store/tunnels'
import Semaphore, { SelectableStatus } from '../semaphore'
import Action from '../action'
interface RowProps {
    tunnel: TunnelState
}

export const Row: FunctionComponent<RowProps> = (
    props: RowProps
) => {
    const t = props.tunnel;
    return (
        <tr>
            <th scope="row" >{t.name}</th>
            <td>{t.bastion}</td>
            <td>{t.address}</td>
            <td>{t.localport}</td>
            <td><Semaphore status={colorFromState(t.state)}/></td>
            <td><Action name={t.name}/></td>
        </tr> 
    )
}

function colorFromState(state : string) : SelectableStatus {   
    switch(state){   
        case 'run' : return SelectableStatus.Running;
        case 'warn': return SelectableStatus.Warning; 
        case 'stop': return SelectableStatus.Off; 
        default: return SelectableStatus.Off;      
    }
}

export default Row