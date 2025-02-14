import { check, sleep } from 'k6';
import http from 'k6/http';
import { randomString, randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export let options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '30s', target: 250 },
    { duration: '30s', target: 500 },
    { duration: '1m', target: 500 },
    { duration: '30s', target: 0 },
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<50'],
  },
};

const BASE_URL = 'http://localhost:8080';
const URLS = new Map();

let resivers = [];

function generateUser(vuId) {
  const prefix = `user_${vuId}_${randomString(5)}`;
  return {
    username: prefix,
    password: randomString(10),
  };
}

export default function () {
  const user = generateUser(__VU);

  resivers.push(user)

  const resiver = resivers[randomIntBetween(0, resivers.length - 1)]

  const authRes = http.post(
    `${BASE_URL}/api/auth`,
    JSON.stringify(user),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(authRes, {
    'Auth status 200': (r) => r.status === 200,
    'Token получен': (r) => r.json().token !== undefined,
  });

  const token = authRes.json('token');

  const headers = {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json',
  };

  sleep(randomIntBetween(1, 3));

  const infoRes = http.get(`${BASE_URL}/api/info`, { headers });
  check(infoRes, {
    'Info status 200': (r) => r.status === 200,
  });

  sleep(Math.random(1, 3));


  URLS.set(0, http.get(`${BASE_URL}/api/buy/powerbank`, { headers }));
  URLS.set(1, http.post(`${BASE_URL}/api/sendCoin`, JSON.stringify({ toUser: resiver.username, amount: 100 }) , { headers }));

  let num = randomIntBetween(0, 1);
  const url = URLS.get(num)
  check(url, {
    'Status 200': (r) => r.status === 200,
  })

  sleep(Math.random(1, 3));

  const final = http.get(`${BASE_URL}/api/info`, { headers });
  check(final, {
    'Info status 200': (r) => r.status === 200,
  });

  sleep(randomIntBetween(1, 5));
}