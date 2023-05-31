import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'

import counterReducer from './features/counter/counterSlice'
import themeReducer from './features/theme/themeSlice'

export function makeStore() {
  return configureStore({
    reducer: { 
      counter: counterReducer,
      theme: themeReducer,
    },
  })
}

const store = makeStore()

export type AppState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  AppState,
  unknown,
  Action<string>
>

export default store
