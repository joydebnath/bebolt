import http from 'k6/http';
import { sleep } from 'k6';

/**
 * TODO: Must install k6 to run the benchmark.
 * Follow the install guide: https://k6.io/docs/get-started/installation/
 */

export const options = {
    scenarios: {
        constant_request_rate: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s', // 1000 iterations per second, i.e. 1000 RPS
            duration: '30s',
            preAllocatedVUs: 100, // how large the initial pool of VUs would be
            maxVUs: 200, // if the preAllocatedVUs are not enough, we can initialize more
        },
    },
};

export default function () {
    http.get('http://localhost:6900');
    sleep(1);
}