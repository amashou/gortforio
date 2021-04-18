import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit"
import { AppThunk, RootState } from "../../app/store"

interface StockState {
    value: [string];
}

const initialState: StockState = {
    value: [""]
}

export const stockSlice = createSlice({
    name: "stock",
    initialState,
    reducers: {
        getStockState: (state, action: PayloadAction<[string]>) => {
            state.value = action.payload;
        }
    }
})

export const { getStockState } = stockSlice.actions;

export const selectStock = { state: RootState } => state.stock.value

export default stockSlice.reducer;