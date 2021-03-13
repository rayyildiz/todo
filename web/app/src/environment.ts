const prod = {
  apiURL: 'https://TBA/api/query',
};

const dev = {
  apiURL: '/api/query',
};

const config = process.env.NODE_ENV === 'development' ? dev : prod;

export const BASE_API = config.apiURL;
