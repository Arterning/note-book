
Create a store

```js
import { createStore, applyMiddleware, compose   } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension'; // 引入 composeWithDevTools 方法
import thunk from 'redux-thunk'; // 引入 redux-thunk

// 初始状态
const initialState = {
    todos: [],
    error: ""
};

// 异步 action
export const fetchTodos = () => {
    return async (dispatch) => {
        try {
            const response = await fetch('https://api.example.com/todos');
            const data = await response.json();
            console.log('fetchTodos', data)
            dispatch({ type: 'FETCH_TODOS_SUCCESS', todos: data });
        } catch (error) {
            dispatch({ type: 'FETCH_TODOS_FAILURE', error: error.message });
        }
    };
};


// Reducer
const todoReducer = (state = initialState, action) => {
    switch (action.type) {
        case 'ADD_TODO':
            return {
                todos: [...state.todos, { id: Date.now(), text: action.text, done: false }]
            };
        case 'TOGGLE_TODO':
            return {
                todos: state.todos.map(todo =>
                    todo.id === action.id ? { ...todo, done: !todo.done } : todo
                )
            };
        case 'FETCH_TODOS_SUCCESS':
            return {
                todos: [...action.todos]
            }
        case 'FETCH_TODOS_FAILURE':
            return {
                todos: [],
                error: action.error
            }
        default:
            return state;
    }
};

// 使用开发者工具调试和中间件
const enhancers = compose(
    applyMiddleware(thunk),
    composeWithDevTools()
);

// 创建 Redux Store
const store = createStore(enhancers);

export default store;

```