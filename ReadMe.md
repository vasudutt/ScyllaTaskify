# ScyllaTaskify

ScyllaTaskify is a simple todo application written in Golang. It allows you to manage your tasks and keep track of your progress.

## Features

-   Add new tasks
-   Mark tasks as completed
-   Delete tasks
-   List all tasks

## Getting Started

Follow these steps to get the application up and running:

1. Clone the repository from GitHub:

    ```bash
    git clone <repository-url>
    ```

    Replace `<repository-url>` with the URL of this repository.

2. Navigate into the cloned repository:

    ```bash
    cd <repository-name>
    ```

    Replace `<repository-name>` with the name of this repository.

3. Run ScyllaDB using Docker:

    ```bash
    docker pull scylladb/scylla
    docker run --name scylladb -p 9042:9042 -d scylladb/scylla
    ```

    This will start a ScyllaDB instance in a Docker container named 'scylladb', and map port 9042 in the container to port 9042 on your local machine.

4. Build and run the application using the Makefile:

    ```bash
    make run
    ```

    This will build the application and start it.

Now, you can use the application to manage your tasks.
