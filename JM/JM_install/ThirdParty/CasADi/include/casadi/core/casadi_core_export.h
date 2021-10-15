
#ifndef CASADI_CORE_EXPORT_H
#define CASADI_CORE_EXPORT_H

#ifdef CASADI_CORE_STATIC_DEFINE
#  define CASADI_CORE_EXPORT
#  define CASADI_CORE_NO_EXPORT
#else
#  ifndef CASADI_CORE_EXPORT
#    ifdef casadi_core_EXPORTS
        /* We are building this library */
#      define CASADI_CORE_EXPORT __attribute__((visibility("default")))
#    else
        /* We are using this library */
#      define CASADI_CORE_EXPORT __attribute__((visibility("default")))
#    endif
#  endif

#  ifndef CASADI_CORE_NO_EXPORT
#    define CASADI_CORE_NO_EXPORT __attribute__((visibility("hidden")))
#  endif
#endif

#ifndef CASADI_CORE_DEPRECATED
#  define CASADI_CORE_DEPRECATED __attribute__ ((__deprecated__))
#endif

#ifndef CASADI_CORE_DEPRECATED_EXPORT
#  define CASADI_CORE_DEPRECATED_EXPORT CASADI_CORE_EXPORT CASADI_CORE_DEPRECATED
#endif

#ifndef CASADI_CORE_DEPRECATED_NO_EXPORT
#  define CASADI_CORE_DEPRECATED_NO_EXPORT CASADI_CORE_NO_EXPORT CASADI_CORE_DEPRECATED
#endif

#define DEFINE_NO_DEPRECATED 0
#if DEFINE_NO_DEPRECATED
# define CASADI_CORE_NO_DEPRECATED
#endif

#endif
