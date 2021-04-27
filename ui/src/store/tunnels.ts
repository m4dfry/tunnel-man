import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { RootState } from "./index";

export interface TunnelState {
    name: string,
    bastion: string,
	address: string,
	localport: string,
    state: 'run' | 'warn' | 'stop'
};

interface TunnelsState {
    tunnels: Array<TunnelState>,
    status: 'empty' | 'pending' | 'success' | 'failed'
};

const initialState: TunnelsState = {
    tunnels: [],
    status: 'empty'
};

export const fetchTunnels = createAsyncThunk(
    'api/tunnels',
    async () => {
        const response = await fetch(`http://localhost:8090/api/tunnels`)
        return response.json()
    }
)

export const tunnelsSlice = createSlice({
    name: 'tunnels',
    initialState,
    // The `reducers` field lets us define reducers and generate associated actions
    reducers: {
        changeTunnelStatus: (state, action) => {
            console.log('changeTunnelStatus', action.payload)
            console.log('changeTunnelStatus', state.tunnels.length)
        }
    },
    // The `extraReducers` field lets the slice handle actions defined elsewhere,
    // including actions generated by createAsyncThunk or in other slices.
    extraReducers: (builder) => {
        builder
            .addCase(fetchTunnels.pending, (state) => {
                state.status = 'pending';
            })
            .addCase(fetchTunnels.fulfilled, (state, action) => {
                state.status = 'success';
                // state.tunnels = action.payload
                var keys = Object.keys(action.payload);
                keys.forEach(function(key) {
                    state.tunnels.push({
                        name: key,
                        bastion: action.payload[key].bastion,
                        address: action.payload[key].address,
                        localport: action.payload[key].localPort,
                        state: 'stop'
                    })
                });
            })
            .addCase(fetchTunnels.rejected, (state) => {
                state.status = 'failed';
            });
    },
});

// Getters
export const getTunnelsList = (state: RootState): TunnelsState =>
  state.tunnels;

export const { changeTunnelStatus } = tunnelsSlice.actions;

export default tunnelsSlice.reducer;