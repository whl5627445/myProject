
#ifndef CASADI_INTEGRATOR_CVODES_EXPORT_H
#define CASADI_INTEGRATOR_CVODES_EXPORT_H

#ifdef CASADI_INTEGRATOR_CVODES_STATIC_DEFINE
#  define CASADI_INTEGRATOR_CVODES_EXPORT
#  define CASADI_INTEGRATOR_CVODES_NO_EXPORT
#else
#  ifndef CASADI_INTEGRATOR_CVODES_EXPORT
#    ifdef casadi_integrator_cvodes_EXPORTS
        /* We are building this library */
#      define CASADI_INTEGRATOR_CVODES_EXPORT __attribute__((visibility("default")))
#    else
        /* We are using this library */
#      define CASADI_INTEGRATOR_CVODES_EXPORT __attribute__((visibility("default")))
#    endif
#  endif

#  ifndef CASADI_INTEGRATOR_CVODES_NO_EXPORT
#    define CASADI_INTEGRATOR_CVODES_NO_EXPORT __attribute__((visibility("hidden")))
#  endif
#endif

#ifndef CASADI_INTEGRATOR_CVODES_DEPRECATED
#  define CASADI_INTEGRATOR_CVODES_DEPRECATED __attribute__ ((__deprecated__))
#endif

#ifndef CASADI_INTEGRATOR_CVODES_DEPRECATED_EXPORT
#  define CASADI_INTEGRATOR_CVODES_DEPRECATED_EXPORT CASADI_INTEGRATOR_CVODES_EXPORT CASADI_INTEGRATOR_CVODES_DEPRECATED
#endif

#ifndef CASADI_INTEGRATOR_CVODES_DEPRECATED_NO_EXPORT
#  define CASADI_INTEGRATOR_CVODES_DEPRECATED_NO_EXPORT CASADI_INTEGRATOR_CVODES_NO_EXPORT CASADI_INTEGRATOR_CVODES_DEPRECATED
#endif

#define DEFINE_NO_DEPRECATED 0
#if DEFINE_NO_DEPRECATED
# define CASADI_INTEGRATOR_CVODES_NO_DEPRECATED
#endif

#endif
