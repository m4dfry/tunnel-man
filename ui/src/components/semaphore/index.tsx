import React from 'react'
import StopLightH from '../../../assets/stop-light-h.svg'
import './index.css'

export enum SelectableStatus {
    Running = "green",
    Warning = "yellow",
    Off = "red"
}

interface ISemaphoreProps {
    status: SelectableStatus
}

const Semaphore = ( {status} : ISemaphoreProps ) => {
    return (
        <div className='semaphore'>
            <div className='semaphore-green'>
                { status == SelectableStatus.Running ? circle('green') : circle('white') }
            </div>
            <div className='semaphore-yellow'>
                { status == SelectableStatus.Warning ? circle('yellow') : circle('white') }
            </div>
            <div className='semaphore-red'>
                { status == SelectableStatus.Off ? circle('red') : circle('white') }
            </div>
        </div>                 
    )
}

const circle = (color : string) => {
    return (
        <svg height="28" width="28">
            <circle cx="16" cy="14" r="8" stroke="black" stroke-width="3" fill={color} />
        </svg>)
}

export default Semaphore