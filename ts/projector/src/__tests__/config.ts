import getProjectorConfig, { Operation } from "../config";

test("should create a print projector config", function () {
  const config = getProjectorConfig({});
  expect(config.operation).toEqual(Operation.Print);
});

test("should create an add projector config", function () {
  const config = getProjectorConfig({
    arguments: ["add", "foo", "bar"],
  });
  expect(config.operation).toEqual(Operation.Add);
  expect(config.args).toEqual(["foo", "bar"]);
});

test("should create a remove projector config", function () {
  const config = getProjectorConfig({
    arguments: ["remove", "foo"],
  });
  expect(config.operation).toEqual(Operation.Remove);
  expect(config.args).toEqual(["foo"]);
});
