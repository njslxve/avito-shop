import { check, sleep } from 'k6';
import http from 'k6/http';
import { randomString, randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export let options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '1m', target: 300 },
    { duration: '30s', target: 0 },
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<50'],
  },
};

const BASE_URL = 'http://localhost:8080';

function generateUser(vuId) {
  const prefix = `user_${vuId}_${randomString(5)}`;
  return {
    username: prefix,
    password: randomString(10),
  };
}

export default function () {
  const user = generateUser(__VU);

  const resivers = []
  resivers.push(user)

  const resiver = resivers[randomIntBetween(0, resivers.length - 1)]

  const authRes = http.post(
    'http://localhost:8080/api/auth',
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

  const responses = http.batch([
    ['GET', `${BASE_URL}/api/info`, null, { headers }],
    ['POST', `${BASE_URL}/api/buy/powerbank`, null , { headers }],
    ['POST', `${BASE_URL}/api/sendCoin`, JSON.stringify({ toUser: resiver.username, amount: 100 }), { headers }],
  ]);

  check(responses[0], { 'Info status 200': (r) => r.status === 200 });
  check(responses[1], { 'Buy status 200': (r) => r.status === 200 });
  check(responses[2], { 'SendCoin status 200': (r) => r.status === 200 });

  sleep(0.3);
}