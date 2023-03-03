import http from 'k6/http';
import { check, group, sleep, fail } from 'k6';

export const options = {
  stages: [
    { duration: '5m', target: 60 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
    { duration: '30m', target: 60 }, // stay at 100 users for 10 minutes
    { duration: '10m', target: 0 }, // ramp-down to 0 users
  ],
  thresholds: {
    'http_req_duration': ['p(99)<1500'], // 99% of requests must complete below 1.5s
    // 'logged in successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s
    // 'logged out successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s

  },
};

const BASE_URL = 'http://localhost:8082';
const EMAIL = 'test@test.com';
const PASSWORD = ':GJ1HNBL:$TT:tf';

export default () => {

  let loginData = { email: EMAIL, password: PASSWORD, };

  let loginRes = http.post(`${BASE_URL}/api/auth/login`, JSON.stringify(loginData), {
    headers: { 'Content-Type': 'application/json' },
  });

  check(loginRes, {
    'logged in successfully': (resp) => resp.json('status') == 'success',
  });


  let logoutRes = http.del(`${BASE_URL}/api/auth/logout`, JSON.stringify(loginData), {
    headers: { 'Content-Type': 'application/json' },
  });
  // console.log(logoutRes.json());

  check(logoutRes, {
    'logged out successfully': (resp) => resp.json('status') == 'success',
  });

  sleep(1);
};
