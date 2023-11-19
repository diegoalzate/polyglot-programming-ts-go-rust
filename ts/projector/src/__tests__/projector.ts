import { Config, Operation } from "../config";
import { Projector } from "../projector";

function getConfig(pwd: string): Config {
  return {
    pwd,
    config: "/foo/bar/baz",
    operation: Operation.Add,
    args: [],
  };
}

function getData() {
  return {
    projector: {
      "/foo/bar/baz/buzz": {
        foo: "bar1",
      },
      "/foo/bar/baz": {
        foo: "bar2",
      },
      "/foo/bar": {
        foo: "bar3",
      },
      "/foo": {
        foo: "bar4",
      },
      "/": {
        foo: "bar5",
        bar: "bazz1",
      },
    },
  };
}

test("getting values", function () {
  const projector = new Projector(getConfig("/foo/bar"), getData());

  expect(projector.getValue("foo")).toEqual("bar3");
  expect(projector.getValue("blaz")).toEqual(undefined);
  expect(projector.getValue("bar")).toEqual("bazz1");
});

test("setting values", function () {
  const projector = new Projector(getConfig("/foo/bar"), getData());

  expect(projector.getValue("foo")).toEqual("bar3");
  projector.setValue("foo", "barNever");
  expect(projector.getValue("foo")).toEqual("barNever");

  const p2 = new Projector(getConfig("/foo"), getData());
  expect(p2.getValue("foo")).toEqual("bar4");

  const p3 = new Projector(getConfig("/foo/bar/baz"), getData());
  expect(p3.getValue("foo")).toEqual("bar2");
});

test("deleting values", function () {
  const projector = new Projector(getConfig("/foo/bar/baz"), getData());

  expect(projector.getValue("foo")).toEqual("bar2");
  projector.deleteValue("foo");
  expect(projector.getValue("foo")).toEqual("bar3");
  projector.deleteValue("foo");
  expect(projector.getValue("foo")).toEqual("bar3");

  const p2 = new Projector(getConfig("/foo/bar"), getData());
  expect(p2.getValue("foo")).toEqual("bar3");
  p2.deleteValue("foo");
  expect(p2.getValue("foo")).toEqual("bar4");
});
