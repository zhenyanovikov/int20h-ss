import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import useClientStore from "../store/clientStore";

function useChangeLanguageEffect() {
  const locale = useClientStore((state) => state.locale);
  const { i18n } = useTranslation();

  useEffect(() => {
    i18n.changeLanguage(locale);
  }, [i18n, locale]);
}

export default useChangeLanguageEffect;
