from locust import HttpUser, TaskSet, task, between

class UserBehavior(TaskSet):
    @task
    def get_network_traffic(self):
        self.client.get("/network-traffic")

class LoadTestUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)