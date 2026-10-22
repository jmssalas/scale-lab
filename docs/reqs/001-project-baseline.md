---
Status: Accepted
Date: 2026-10-19
---

# REQ 001: Project Baseline

## Description

Design and build an event-processing service that allows external devices or services to send their events in real time, so that they can be processed in a decoupled manner.

## Functional Requirements

- The service should provide a endpoint as a single entry point to receive the events.
- The service should process the events by just storing them as they arrived, keeping the received timestamp.
- The service should be prepared to process the events according to their type.
- There should be an external service —`load generator`— capable of simulating external devices or service requests.
- The `load generator` service should make request according to a configurable rate.

## Non-Functional Requirements

- The service should be as simple as possible.
- The service should be the baseline for the next requirements and iterations.
- There is no requirement regarding availability, low latency or scalability.
- The following metrics should be available from the main service:
    - total requests, request time and processing time.
- The following metrics should be available from the load generator service:
    - total requests, total dropped requests, target requests rate and client-side latency per request.
