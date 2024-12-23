---
title: Analytics Dashboard
date: 2024-02-20
template: dashboard
refresh_rate: 30
data_sources:
  - type: metrics
    endpoint: /api/metrics
  - type: events
    endpoint: /api/events
layout:
  columns: 3
  rows: 2
---

<div class="dashboard-grid" style="display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px;">
  <!-- Metrics Panel -->
  <div class="metric-card" style="background: linear-gradient(135deg, #00C9FF 0%, #92FE9D 100%); padding: 1em; border-radius: 10px;">
    <h3 style="color: white;">Active Users</h3>
    {{with .metrics.active_users}}
      <div class="metric-value">{{.count}}</div>
      <div class="metric-trend {{if gt .growth 0}}positive{{else}}negative{{end}}">
        {{.growth}}% from last week
      </div>
    {{end}}
  </div>

  <!-- Performance Panel -->
  <div class="performance-card" style="background: linear-gradient(135deg, #FF6B6B 0%, #FFE66D 100%); padding: 1em; border-radius: 10px;">
    <h3 style="color: white;">System Performance</h3>
    <div class="performance-metrics">
      <div class="cpu-usage">CPU: {{.metrics.cpu}}%</div>
      <div class="memory-usage">Memory: {{.metrics.memory}}%</div>
    </div>
  </div>

  <!-- Events Panel -->
  <div class="events-card" style="background: linear-gradient(135deg, #7F7FD5 0%, #91EAE4 100%); padding: 1em; border-radius: 10px;">
    <h3 style="color: white;">Recent Events</h3>
    <div class="events-list">
      {{range .events}}
        <div class="event-item">
          <span class="event-type">{{.type}}</span>
          <span class="event-time">{{.timestamp}}</span>
        </div>
      {{end}}
    </div>
  </div>
</div>

## Custom CSS Variables

```css
:root {
  --primary-gradient: linear-gradient(135deg, #00C9FF 0%, #92FE9D 100%);
  --secondary-gradient: linear-gradient(135deg, #FF6B6B 0%, #FFE66D 100%);
  --tertiary-gradient: linear-gradient(135deg, #7F7FD5 0%, #91EAE4 100%);
  --card-border-radius: 10px;
  --card-padding: 1em;
  --text-color-light: white;
  --grid-gap: 20px;
}
```

## JavaScript Components

```javascript
// Refresh dashboard data
function refreshDashboard() {
  const refreshRate = {{.refresh_rate}} * 1000;
  setInterval(() => {
    fetchMetrics();
    fetchEvents();
  }, refreshRate);
}
```