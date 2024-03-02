import { useQueryClient, useQuery, useMutation } from "@tanstack/react-query";
import { enqueueSnackbar } from "notistack";
import { useTranslation } from "react-i18next";
import apiClient from "./apiClient";
import { TEACHERS_URL, TEACHERS_KEY } from "../constants/teacher";

async function getTeachers() {
  const res = await apiClient.get(TEACHERS_URL);
  return res.data;
}

async function inviteTeacher(teacher) {
  const res = await apiClient.post(TEACHERS_URL, teacher);
  return res.data;
}

export function useGetTeachers() {
  return useQuery({
    queryKey: [TEACHERS_KEY],
    queryFn: getTeachers,
  });
}

export function useInviteTeacher() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: inviteTeacher,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.inviteTeacher.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [TEACHERS_KEY] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.inviteTeacher.error"), { variant: "error" });
    },
  });
}
