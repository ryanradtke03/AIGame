import { lazy } from "react";

const pages = import.meta.glob("../pages/**/index.tsx");

const routes = Object.entries(pages).map(([path, resolver]) => {
  const match = path.match(/\.\/pages\/(.+)\/index\.tsx/);
  const routePath = match?.[1]?.toLocaleLowerCase() || "/";
  return {
    path: routePath === "landing" ? "/" : `/${routePath}`,
    element: lazy(resolver as () => Promise<{ default: React.ComponentType }>),
  };
});

export default routes;
