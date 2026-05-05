# GolangTasks

A comprehensive collection of Go programming tasks, algorithms, data structures, concurrency patterns, and interview preparation exercises.

## Project Info

- **Module:** `github.com/stepan41k/GolangTasks`
- **Go Version:** 1.25.3

## Directory Structure

### Algorithms

Classic algorithm implementations:

| Directory | Description |
|---|---|
| `algorithms/backtracking/` | Backtracking algorithm |
| `algorithms/dijkstra/` | Dijkstra's shortest-path algorithm |
| `algorithms/fibonachi/` | Fibonacci number computation |
| `algorithms/kmp/` | Knuth-Morris-Pratt string matching |
| `algorithms/quick_sort/` | Quick sort algorithm |
| `algorithms/token_bucket/` | Token bucket rate limiting |

### Async/Await

| Directory | Description |
|---|---|
| `async-await/` | Future/Promise pattern emulation using generics |

### Bit Manipulation

| Directory | Description |
|---|---|
| `bit_tasks/check_bit/` | Check if a specific bit is set |
| `bit_tasks/find_unique/` | Find unique element using XOR |
| `bit_tasks/is_power_of_two/` | Check if a number is a power of two |
| `bit_tasks/kernigan_count_set_bits/` | Count set bits (Kernighan's algorithm) |
| `bit_tasks/set_bit/` | Set a specific bit |
| `bit_tasks/toggle_bit/` | Toggle a specific bit |
| `bit_tasks/zero_bit/` | Clear a specific bit |

### Concurrency Patterns

| Directory | Description |
|---|---|
| `concurency_patterns/batch_processing/` | Batch processing with goroutines |
| `concurency_patterns/cache_implementation/` | Concurrent-safe cache |
| `concurency_patterns/distributed_query/` | Distributed query pattern |
| `concurency_patterns/fan_in_fan_out/` | Fan-in / Fan-out channel pattern |
| `concurency_patterns/merge_channels/` | Merging multiple channels |
| `concurency_patterns/pipeline/` | Pipeline pattern |
| `concurency_patterns/semaphore/` | Semaphore using buffered channels |
| `concurency_patterns/worker_pool/` | Worker pool implementation |

### Concurrency Primitives

| Directory | Description |
|---|---|
| `concurrency/atomic_select_example/` | `sync/atomic` + `select` usage |
| `concurrency/context_with_timeout/` | Context with cancellation and timeout |
| `concurrency/long_calculation/` | Cancellable long-running computation |
| `concurrency/ticker_example/` | `time.Ticker` periodic execution |
| `concurrency/wait_group/` | `sync.WaitGroup` usage |

### Data Structures

**Arrays:** `dynamic_array/`, `matrix/`, `statick_array/`

**Binary Heap:** `max/`, `min/`

**Graphs:** `acyclic/`, `bipartite/`, `cyclic/`, `direct/`, `strongly_connected/`, `tree/`, `undirect/`, `weighted/`

**Hash Table:** `hash_table/`

**Linked Lists:** `circle/`, `double/`, `single/`

**Queues:** `deque/`, `fifo/`, `priority/`

**Stack:** `stack/`

**Trees:** `avl/`, `b/`, `b_plus/`, `binary/`, `is_valid/`, `prefix/`, `red-black/`

### Go Language Features

| Directory | Description |
|---|---|
| `golangs_features/batch_proccessing/` | Batch processing (v1) |
| `golangs_features/batch_proccessing_v2/` | Batch processing (v2) |
| `golangs_features/batch_proccessing_v3/` | Batch processing (v3) |
| `golangs_features/closure/` | Closures in Go |
| `golangs_features/concurrent_safe_map/` | Thread-safe map |
| `golangs_features/dependency_injection/` | Dependency injection pattern |
| `golangs_features/function/` | First-class functions |
| `golangs_features/gracefull_shutdown/` | Graceful server shutdown |
| `golangs_features/hard_batch_proccessing/` | Advanced batch processing |
| `golangs_features/rate_limiter/` | Rate limiting |
| `golangs_features/resource_pool/` | Resource/connection pool |
| `golangs_features/ttl_cache/` | TTL-based cache |
| `golangs_features/urls_fetching_err_group/` | Concurrent URL fetching with errgroup |
| `golangs_features/worker_pool/` | Worker pool pattern |

### Interview Tasks

| Directory | Description |
|---|---|
| `interview_tasks/err_group/` | errgroup usage |
| `interview_tasks/golang_gemini_interview/` | AI-assisted interview sessions (organized by date) |
| `interview_tasks/interview_popular_tasks/` | Common interview problems |
| `interview_tasks/sber/` | Sber interview tasks |
| `interview_tasks/worker_pool/` | Worker pool interview task |
| `interview_tasks/yandex/` | Yandex interview tasks |

### Linux

Shell scripting and sysadmin tasks:

| Directory | Description |
|---|---|
| `linux/bash/iterate_ips/` | IP iteration script |
| `linux/commands.sh` | Shell commands reference |
| `linux/find/find_file_condition/` | Find files by condition |
| `linux/grep/black_box/` | Black box grep task |
| `linux/grep/logs_analysis/` | Log analysis with grep |
| `linux/kuber_monitoring/kuber_resource_limits/` | Kubernetes resource monitoring |
| `linux/memory/inodes/` | Inode inspection |
| `linux/netcat/network_availability/` | Network availability check |
| `linux/ss/find_port/` | Find listening ports |

### SQL Tasks

| Directory | Description |
|---|---|
| `sql_tasks/avg_categories_price/` | Average price by category |
| `sql_tasks/rank_task/` | SQL RANK / window functions |
| `sql_tasks/vip_clients/` | VIP client identification |

### Standard Library

| Directory | Description |
|---|---|
| `std/reflect/type/` | `reflect.Type` usage |
| `std/reflect/value/` | `reflect.Value` usage |

### Testing

| Directory | Description |
|---|---|
| `testing/benchmark_simple_task/` | Simple benchmark example |

### Practice Tasks

| Directory | Description |
|---|---|
| `zadachki_41k/dedupliceate_slice/` | Slice deduplication |
| `zadachki_41k/gracefully/` | Graceful shutdown |
| `zadachki_41k/is_anagramm/` | Anagram checker |
| `zadachki_41k/is_palindrome/` | Palindrome checker |
| `zadachki_41k/len_same_symbols/` | Count matching symbols |
| `zadachki_41k/parallel_squares/` | Parallel computation |

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/stretchr/testify` | Testing framework |
| `golang.org/x/sync` | Advanced concurrency (errgroup, semaphore) |
| `golang.org/x/exp` | Experimental Go packages |

## License

See [LICENSE](LICENSE) for details.
