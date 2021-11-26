<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use TheSeer\Tokenizer\Token;

class AuthController extends Controller
{
    public function login(Request $request): JsonResponse
    {
        $name = $request->get('name');

        if (User::query()->where('name', '=', $name)->exists()) {
            $token = (new Token(1, $name, $name));
            session($name, $token);
        } else {
            $token = 'None';
        }

        return response()
            ->json([
                'token' => $token
            ]);
    }
}
