import React, { FunctionComponent, useEffect} from 'react'
import './App.css'
import "bootstrap/dist/css/bootstrap.css"; // Import precompiled Bootstrap css
import Table from './components/table';
import { useAppDispatch, useAppSelector } from './hooks';
import { fetchTunnels, getTunnelsList, changeTunnelStatus } from './store/tunnels';


export const App: FunctionComponent = () => {
  const dispatch = useAppDispatch();
  const tunnelsList = useAppSelector(getTunnelsList);

  useEffect(() => {
    dispatch(fetchTunnels());

    var ws = new WebSocket('ws://localhost:8090/api/ws');
    ws.onmessage = function(event) {
      dispatch(changeTunnelStatus(event.data));
    }
    
  }, []);
  
  return (
    <div className="App">
      <div className="app-layout">
        <div className="header box"><h1>Tunnel🤦‍♂️Man</h1></div>
        <div className="table box">
          {tunnelsList.status === 'success' && <Table tunnels={tunnelsList.tunnels}/> }
        </div>
        <div className="input box">
          <input type="text" placeholder=" 🔎 Filter ..."/>
        </div>
      </div>
    </div>
  );
}

export default App;
