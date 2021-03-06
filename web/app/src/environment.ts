const prod = {
  apiURL: 'https://api-todo.rayyildiz.dev/api/query',
  swEnable: false,
};

const dev = {
  apiURL: '/api/query',
  swEnable: false,
};

const config = process.env.NODE_ENV === 'development' ? dev : prod;

export const BASE_API = config.apiURL;
export const ENABLE_SW = config.swEnable;
export const LOCALSTORAGE_AUTH_KEY = "auth:key";
