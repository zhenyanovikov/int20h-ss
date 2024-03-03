import useClientStore from "../store/clientStore";

function useIsAuthenticated() {
  const token = useClientStore((state) => state.token);

  return !!token;
}

export default useIsAuthenticated;
