import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    // stages: [
    //     { duration: '10s', target: 5 },
    //     { duration: '1m', target: 20 },
    //     { duration: '10', target: 0 },
    // ],
    stages: [
        { duration: '1', target: 1 },
    ],
};

export default function () {
    const res = http.get('http://127.0.0.1:1323/v1/health/ping');
    check(res, { 'status was 200': (r) => r.status == 200 });
    sleep(1);
}
