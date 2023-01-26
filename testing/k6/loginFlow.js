import http from 'k6/http';
import { check, group, sleep, fail } from 'k6';

export const options = {
  vus: 1, // 1 user looping for 1 minute
  duration: '1m',

  thresholds: {
    http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
  },
};

const BASE_URL = 'http://localhost:8082';
const EMAIL = 'test@test.com';
const PASSWORD = 'testTest@Test';

export default () => {

  let loginData = { email: EMAIL, password: PASSWORD, };

  let loginRes = http.post(`${BASE_URL}/auth/login`, JSON.stringify(loginData), {
    headers: { 'Content-Type': 'application/json' },
  });

  check(loginRes, {
    'logged in successfully': (resp) => resp.json('status') == 'success',
  });


  let logoutRes = http.del(`${BASE_URL}/auth/logout`, JSON.stringify(loginData), {
    headers: { 'Content-Type': 'application/json' },
  });
  // console.log(logoutRes.json());

  check(logoutRes, {
    'logged out successfully': (resp) => resp.json('status') == 'success',
  });



  sleep(.1);
};
