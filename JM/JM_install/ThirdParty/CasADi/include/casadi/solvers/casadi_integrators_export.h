
#ifndef CASADI_INTEGRATORS_EXPORT_H
#define CASADI_INTEGRATORS_EXPORT_H

#ifdef CASADI_INTEGRATORS_STATIC_DEFINE
#  define CASADI_INTEGRATORS_EXPORT
#  define CASADI_INTEGRATORS_NO_EXPORT
#else
#  ifndef CASADI_INTEGRATORS_EXPORT
#    ifdef casadi_integrators_EXPORTS
        /* We are building this library */
#      define CASADI_INTEGRATORS_EXPORT __attribute__((visibility("default")))
#    else
        /* We are using this library */
#      define CASADI_INTEGRATORS_EXPORT __attribute__((visibility("default")))
#    endif
#  endif

#  ifndef CASADI_INTEGRATORS_NO_EXPORT
#    define CASADI_INTEGRATORS_NO_EXPORT __attribute__((visibility("hidden")))
#  endif
#endif

#ifndef CASADI_INTEGRATORS_DEPRECATED
#  define CASADI_INTEGRATORS_DEPRECATED __attribute__ ((__deprecated__))
#endif

#ifndef CASADI_INTEGRATORS_DEPRECATED_EXPORT
#  define CASADI_INTEGRATORS_DEPRECATED_EXPORT CASADI_INTEGRATORS_EXPORT CASADI_INTEGRATORS_DEPRECATED
#endif

#ifndef CASADI_INTEGRATORS_DEPRECATED_NO_EXPORT
#  define CASADI_INTEGRATORS_DEPRECATED_NO_EXPORT CASADI_INTEGRATORS_NO_EXPORT CASADI_INTEGRATORS_DEPRECATED
#endif

#define DEFINE_NO_DEPRECATED 0
#if DEFINE_NO_DEPRECATED
# define CASADI_INTEGRATORS_NO_DEPRECATED
#endif

#endif
