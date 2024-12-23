---
title: Interactive Coding Tutorial
date: 2024-02-15
template: interactive
theme:
  code-background: "#1E1E1E"
  code-text: "#D4D4D4"
  highlight: "#569CD6"
interactive_elements:
  - type: code-editor
    language: python
    initial_content: |
      def greet(name):
          return f"Hello, {name}!"
  - type: quiz
    questions: 3
---

<div class="tutorial-container">
  <header style="background: linear-gradient(45deg, #2C3E50, #3498DB);">
    <h1>Learn Python: Day 1</h1>
  </header>

  <div class="lesson-section">
    {{range .interactive_elements}}
      <div class="interactive-block" 
           style="background-color: {{if eq .type "code-editor"}}#1E1E1E{{else}}#F8F9FA{{end}}">

        {{if eq .type "code-editor"}}
          <pre><code class="language-python">{{.initial_content}}</code></pre>
        {{end}}
        
        {{if eq .type "quiz"}}
          <div class="quiz-container">
            <!-- Quiz template goes here -->
          </div>
        {{end}}
      </div>
    {{end}}
  </div>
</div>

## Code Examples

```python
# Example 1: Basic Function
def calculate_area(radius):
    return 3.14 * radius * radius

# Example 2: List Comprehension
squares = [x**2 for x in range(10)]
```

---