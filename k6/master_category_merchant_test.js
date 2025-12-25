import http from 'k6/http';
import { check, sleep } from 'k6';

// export const options = {
//   vus: 1000,          // virtual users (banyaknya user bersamaan)
//   duration: '30s',  // durasi tes
// };

export const options = {
  stages: [
    { duration: '10s', target: 50 },
    { duration: '20s', target: 200 },
    { duration: '30s', target: 500 },
    { duration: '20s', target: 200 },
    { duration: '10s', target: 0 },
  ],
};

export default function () {
  let res = http.get('http://127.0.0.1:3001/merchants/categories');
  
  check(res, {
    'status 200': (r) => r.status === 200,
  });
  
  sleep(1);
}
