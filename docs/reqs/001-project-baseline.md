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

## Non-Functional Requirements

- The service should be as simple as possible.
- The service should be the baseline for the next requirements and iterations.
- There is no requirement regarding availability, low latency or scalability.
- The following metrics should be available: total requests, request time and processing time.
