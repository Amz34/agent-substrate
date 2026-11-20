import unittest
from pathlib import Path
import tempfile
import shutil
from substrate import SubstrateMesh, AXTask, AXWorkspace, AXGateway, AXModel

class TestSubstrateMesh(unittest.TestCase):
    def setUp(self):
        self.test_dir = Path(tempfile.mkdtemp())
        self.mesh = SubstrateMesh(self.test_dir)

    def tearDown(self):
        shutil.rmtree(self.test_dir, ignore_errors=True)

    def test_hibernate_and_wake(self):
        state = {"step": 3, "registers": {"pc": 1024}, "scratch": "active_buffer"}
        res = self.mesh.hibernate("test-agent", state)
        self.assertEqual(res["status"], "SNAPSHOT_COMPACTED")
        self.assertLess(res["latency_ms"], 50.0)

        wake_res = self.mesh.wake("test-agent")
        self.assertIsNotNone(wake_res)
        self.assertEqual(wake_res["status"], "RESTORED")
        self.assertLess(wake_res["latency_ms"], 500.0)
        self.assertEqual(wake_res["state"]["step"], 3)

    def test_declarative_dispatch(self):
        task = AXTask(name="unit_task", workspace_ref="/tmp", gateway_ref="strict", model_ref="gemini")
        ws = AXWorkspace(path="/tmp")
        gw = AXGateway(allowed_domains=["api.github.com"])
        mdl = AXModel(provider="google", model_name="gemini-3.8-flash")

        dispatch_res = self.mesh.dispatch_declarative(task, ws, gw, mdl)
        self.assertEqual(dispatch_res["status"], "INITIALIZED")
        self.assertTrue(Path(dispatch_res["sandbox_path"]).exists())

if __name__ == "__main__":
    unittest.main()
