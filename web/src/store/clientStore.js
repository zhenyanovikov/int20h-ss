import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { STORE_NAME } from "../constants/store";
import { LOCALE } from "../constants/localisation";
import { THEME_MODE } from "../constants/theme";

const useClientStore = create(
  persist(
    (set) => ({
      themeMode: THEME_MODE.LIGHT,
      locale: LOCALE.UK,
      toggleThemeMode: () => {
        set((state) => ({
          themeMode:
            state.themeMode === THEME_MODE.DARK
              ? THEME_MODE.LIGHT
              : THEME_MODE.DARK,
        }));
      },
      setLocale: (locale) => {
        set({ locale });
      },
    }),
    {
      name: STORE_NAME,
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useClientStore;
