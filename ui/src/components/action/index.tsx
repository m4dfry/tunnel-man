import React from 'react'
import './index.css'

const icons = {
    play:"M0 0v6l6-3-6-3z",
    stop:"M0 0v6h6v-6h-6z"
}

function DrawIcon(icon: string) {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16">
            <path d={icon} transform="translate(1 1) scale(2 2)" />
        </svg>
    )
}

function Action() {
    return (
        <div className='action'>
            <button type="button" className="btn btn-outline-success btn-sm">
               { DrawIcon(icons.play) }
            </button>
            <button type="button" className="btn btn-outline-danger btn-sm">
                { DrawIcon(icons.stop) }
            </button>
        </div>    
    )
}



export default Action